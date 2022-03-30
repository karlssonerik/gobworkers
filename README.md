# Gobworkers

This is small project to try and save time from rewriting go channel workers.

## Basic flow
 1. Create a worker pool with desired number of works with `gobworkers.New`
 2. Add work with `WorkerPool.AddWork`
 3. Wait for the workers to finish with `WorkerPool.WaitForWorkToBeDone`

## Example
See `example/main.go`