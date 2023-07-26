#!/bin/bash

node_config_file_name="node-config.json"
node_output_file="bench-output.json"

dir="$1"
hosts="$2"
settings="$3"

num_hosts=$(wc -l < "$hosts")

mkdir "$dir"
cp "$settings" "$dir/parameter-set.yaml" || exit 1
cp "$hosts" "$dir/hosts" || exit 1

python3 scripts/hosts-to-membership.py < "$dir/hosts" > "$dir/membership.json" || exit 1
go run ../cmd/bench params -m "$dir/membership.json" -o "$node_config_file_name" -s "$settings" -d "$dir" -w 4 || exit 1

ansible-playbook -i "$dir/hosts" setup.yaml

for exp_dir in "$dir"/[0-9][0-9][0-9][0-9]; do
  exp_id=$(basename "$exp_dir")
  ansible-playbook -i "$dir/hosts" --forks "$num_hosts" --extra-vars "node_config_file=$exp_dir/$node_config_file_name exp_id=$exp_id output_dir='$exp_dir' node_output_file=$node_output_file" run-benchmark.yaml || exit 1
  python3 scripts/analyze-bench-output.py "$exp_id" "$exp_dir/$node_config_file_name" "$exp_dir/results.json" "$exp_dir"/data/*$node_output_file || exit 1
done

python3 scripts/aggregate-data.py "$dir"/all-results.csv "$dir"/[0-9][0-9][0-9][0-9]/results.json
