task :build do
  sh 'go build'
end

task :install do
  sh 'go get'
end

task :clean do
  sh 'git clean -dfx'
end
