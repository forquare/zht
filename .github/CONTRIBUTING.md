# Contributing

## Issues

Issues are very valuable to this project:
* Ideas are a valuable source of contributions that anyone can make
* Problems show where this project is lacking
* With a question you show where contributors can improve the user experience

Thank you for creating them.

## Pull Requests

Pull Requests are a great way to get your ideas into this repository.

When deciding whether a Pull Request will be merged, the following is considered:

* Does it state intent?  
  You should be clear which problem you're trying to solve with your contribution.

  For example:
  
  > Fix search
  
  Doesn't describe what is being fixed or what was broken, whereas:
  
  > Fix search to actually use POSIX regex when selected
  > 
  > After a recent update the `-R` flag for search stopped using POSIX regex and instead was using the default Go regex, this commit reverts it to using POSIX regex as per the arguments help description.
  
  Describes the problem that has been found, and the Pull Request shows the action taken to solve it.

* Applies good Git history management
  
  * Pull Requests should be up-to-date with the main branch.
  * Each Pull Request should do one thing.
  * Commits should be atomic (i.e. do only one thing), most Pull Requests will contain a single commit with a descriptive message similar to the example given above.
  * Larger Pull Requests might contain multiple commits, each should have a descriptive message and not depend on each other
    * For example, if _one_ of the commits was reverted the entire PR wouldn't break functionality.

* Does it follow the contributor covenant?
  This repository has a [code of conduct](.github/CODE_OF_CONDUCT.md), if a Pull Request does not respect it then it will be removed.

There are always exceptions to these guidelines, and it is down to the maintainers discretion when to make those exceptions.