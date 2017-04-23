Pod::Spec.new do |spec|
  spec.name         = 'Gele'
  spec.version      = '{{.Version}}'
  spec.license      = { :type => 'GNU Lesser General Public License, Version 3.0' }
  spec.homepage     = 'https://github.com/elementrem/go-elementrem'
  spec.authors      = { {{range .Contributors}}
		'{{.Name}}' => '{{.Email}}',{{end}}
	}
  spec.summary      = 'iOS Elementrem Client'
  spec.source       = { :git => 'https://github.com/elementrem/go-elementrem.git', :commit => '{{.Commit}}' }

	spec.platform = :ios
  spec.ios.deployment_target  = '9.0'
	spec.ios.vendored_frameworks = 'Frameworks/Gele.framework'

	spec.prepare_command = <<-CMD
    curl https://gelestore.blob.core.windows.net/builds/{{.Archive}}.tar.gz | tar -xvz
    mkdir Frameworks
    mv {{.Archive}}/Gele.framework Frameworks
    rm -rf {{.Archive}}
  CMD
end
