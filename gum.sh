#!/bin/bash

gum style --border normal --margin "1" --padding "1 2" --border-foreground 212 "Hello! Welcom to $(gum style --foreground 221 'Gum')."

NAME=$(gum input --placeholder "What is your name?")

echo -e "I thought that was you $(gum style --foreground 212 "$NAME")."
