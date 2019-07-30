# git remove file in a commit
I think other answers here are wrong, because this is a question of moving the mistakenly committed files back to the staging area from the previous commit, without cancelling the changes done to them. This can be done like Paritosh Singh suggested:

```
git reset --soft HEAD^
```

or

```
git reset --soft HEAD~1
```

Then reset the unwanted files in order to leave them out from the commit:

```
git reset HEAD path/to/unwanted_file
```

Now commit again, you can even re-use the same commit message:

```
git commit -c ORIG_HEAD  
```
