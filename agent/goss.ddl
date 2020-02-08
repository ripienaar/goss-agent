metadata :name        => "goss",
         :description => "System validation using Goss",
         :author      => "R.I.Pienaar <rip@devco.net>",
         :license     => "Apache-2",
         :version     => "0.0.1",
         :url         => "https://devco.net",
         :provider    => "external",
         :timeout     => 20


action "validate", :description => "Validate the system" do
  display :failed

  input :sleep,
        :prompt      => "Sleep Duration",
        :description => "Time to sleep between retries when",
        :type        => :string,
        :default     => "1s",
        :validation  => '\d+[hms]',
        :maxlength   => 3,
        :optional    => true


  input :gossfile,
        :prompt      => "Goss file",
        :description => "Path to the gossfile or it's contents as YAML/JSON",
        :type        => :string,
        :validation  => '.+',
        :maxlength   => 10240,
        :optional    => false


  input :max_concurrency,
        :prompt      => "Maximum Concurrency",
        :description => "Max number of tests to run concurrently",
        :type        => :integer,
        :default     => 50,
        :optional    => true

  input :package,
        :prompt      => "Package Type",
        :description => "The type of package manager to use",
        :type        => :list,
        :list        => ["rpm", " deb", " apk", " pacman"],
        :optional    => true


  input :retry_timeout,
        :prompt      => "Retry Timeout",
        :description => "Retry on failure so long as elapsed + sleep time is less than this",
        :type        => :string,
        :validation  => '\d+[hms]',
        :maxlength   => 3,
        :optional    => true


  input :vars,
        :prompt      => "Variables",
        :description => "Path to the variables or it's contents as YAML/JSON",
        :type        => :string,
        :validation  => '.+',
        :maxlength   => 10240,
        :optional    => true

  output :code,
         :description => "Exit Code",
         :default     => 1,
         :type        => "integer",
         :display_as  => "Exit Code"

  output :result,
         :description => "Output Result as JSON",
         :default     => "{}",
         :type        => "string",
         :display_as  => "Output"

end

