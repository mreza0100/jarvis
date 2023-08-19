# boot:
# 	go build .

# 	sudo rm -rf /home/noroot/exec
# 	sudo mkdir -p /home/noroot/exec
# 	sudo mv ./gptjarvis /home/noroot/exec/

# 	sudo -u noroot /home/noroot/exec/gptjarvis boot

boot:
	go run . boot
