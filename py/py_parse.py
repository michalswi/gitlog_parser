#!/usr/bin/env python

import git
from git import Repo
import json
import argparse


MAIN_DICT = {
    "data": [],
    "status": "completed",
    "error": "null",
}

NUM_OF_COMMITS = 5

ap = argparse.ArgumentParser()
ap.add_argument("-p", "--path", 
                required = True, 
                help = "provide path to your repo, e.g. '/tmp/repo'")
args = vars(ap.parse_args())

REPO = Repo(args["path"])


# example, display all commits
# g = git.Git(REPO) 
# # loginfo = g.log('--since=2013-09-01','--author=...','--pretty=tformat:','--numstat')
# loginfo = g.log()
# print(loginfo)


def upload_commit(commit):
    # NO -> commit.merge
    tmp_dict = {}
    tmp_dict["commit"] = commit.hexsha
    tmp_dict["author"] = commit.author.name
    tmp_dict["date"] = str(commit.authored_datetime)
    tmp_dict["message"] = commit.summary
    # print(str("count: {} and size: {}".format(commit.count(), commit.size)))
    return tmp_dict


def print_repo_details(repo):
    print('Repo description: {}'.format(repo.description))
    print('Repo active branch is {}'.format(repo.active_branch))
    for remote in repo.remotes:
        print('Remote named "{}" with URL "{}"'.format(remote, remote.url))
    print('Last commit for repo is {}.'.format(str(repo.head.commit.hexsha)))


def main(main_dict, repo):
    if not repo.bare:
        # print('Repo at {} successfully loaded.'.format(args["path"]))
        # print_repo_details(repo)
        
        # limit the number of commits
        commits = list(repo.iter_commits('master'))[:NUM_OF_COMMITS]
        # commits = list(repo.iter_commits('master'))[:]
        for commit in commits:
            main_dict["data"].append(upload_commit(commit))
    return json.dumps(main_dict)


if __name__ == "__main__":
    print(main(MAIN_DICT, REPO))
