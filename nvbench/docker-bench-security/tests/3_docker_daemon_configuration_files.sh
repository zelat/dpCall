#!/bin/bash

check_3() {
  logit ""
  local id="3"
  local desc="Docker daemon configuration files"
  checkHeader="$id - $desc"
  info "$checkHeader"
  startsectionjson "$id" "$desc"
}

check_3_1() {
  local id="3.1"
  local desc="确保docker.service文件的所有权为root:root (Automated)"
  local remediation="Find out the file location: systemctl show -p FragmentPath docker.service. If the file does not exist, this recommendation is not applicable. If the file does exist, you should run the command chown root:root <path>, in order to set the ownership and group ownership for the file to root."
  local remediationImpact="None."
  local check="$id - $desc"
  starttestjson "$id" "$desc"

  file=$(get_service_file docker.service)
  if [ -f "$file" ]; then
    if [ "$(stat -c %u%g "$file")" -eq 00 ]; then
      pass -s "$check"
      logcheckresult "PASS"
      return
    fi
    warn -s "$check"
    warn "     * Wrong ownership for $file"
    logcheckresult "WARN" "Wrong ownership for $file"
    return
  fi
  info -c "$check"
  info "     * File not found"
  logcheckresult "INFO" "File not found"
}

check_3_2() {
  local id="3.2"
  local desc="确保docker.service文件权限被合理设置 (Automated)"
  local remediation="Find out the file location: systemctl show -p FragmentPath docker.service. If the file does not exist, this recommendation is not applicable. If the file exists, run the command chmod 644 <path> to set the file permissions to 644."
  local remediationImpact="None."
  local check="$id - $desc"
  starttestjson "$id" "$desc"

  file=$(get_service_file docker.service)
  if [ -f "$file" ]; then
    if [ "$(stat -c %a "$file")" -le 644 ]; then
      pass -s "$check"
      logcheckresult "PASS"
      return
    fi
    warn -s "$check"
    warn "     * Wrong permissions for $file"
    logcheckresult "WARN" "Wrong permissions for $file"
    return
  fi
  info -c "$check"
  info "     * File not found"
  logcheckresult "INFO" "File not found"
}

check_3_3() {
  local id="3.3"
  local desc="确保docker.socket文件所有权为root:root (Automated)"
  local remediation="Find out the file location: systemctl show -p FragmentPath docker.socket. If the file does not exist, this recommendation is not applicable. If the file exists, run the command chown root:root <path> to set the ownership and group ownership for the file to root."
  local remediationImpact="None."
  local check="$id - $desc"
  starttestjson "$id" "$desc"

  file=$(get_service_file docker.socket)
  if [ -f "$file" ]; then
    if [ "$(stat -c %u%g "$file")" -eq 00 ]; then
      pass -s "$check"
      logcheckresult "PASS"
      return
    fi
    warn -s "$check"
    warn "     * Wrong ownership for $file"
    logcheckresult "WARN" "Wrong ownership for $file"
    return
  fi
  info -c "$check"
  info "     * File not found"
  logcheckresult "INFO" "File not found"
}

check_3_4() {
  local id="3.4"
  local desc="确保docker.socket文件权限被设置为644或更多限制性 (Automated)"
  local remediation="Find out the file location: systemctl show -p FragmentPath docker.socket. If the file does not exist, this recommendation is not applicable. If the file does exist, you should run the command chmod 644 <path> to set the file permissions to 644."
  local remediationImpact="None."
  local check="$id - $desc"
  starttestjson "$id" "$desc"

  file=$(get_service_file docker.socket)
  if [ -f "$file" ]; then
    if [ "$(stat -c %a "$file")" -le 644 ]; then
      pass -s "$check"
      logcheckresult "PASS"
      return
    fi
    warn -s "$check"
    warn "     * Wrong permissions for $file"
    logcheckresult "WARN" "Wrong permissions for $file"
    return
  fi
  info -c "$check"
  info "     * File not found"
  logcheckresult "INFO" "File not found"
}

check_3_5() {
  local id="3.5"
  local desc="确保/etc/docker的目录所有权为root:root (Automated)"
  local remediation="You should run the following command: chown root:root /etc/docker. This sets the ownership and group ownership for the directory to root."
  local remediationImpact="None."
  local check="$id - $desc"
  starttestjson "$id" "$desc"

  directory="/etc/docker"
  if [ -d "$directory" ]; then
    if [ "$(stat -c %u%g $directory)" -eq 00 ]; then
      pass -s "$check"
      logcheckresult "PASS"
      return
    fi
    warn -s "$check"
    warn "     * Wrong ownership for $directory"
    logcheckresult "WARN" "Wrong ownership for $directory"
    return
  fi
  info -c "$check"
  info "     * Directory not found"
  logcheckresult "INFO" "Directory not found"
}

check_3_6() {
  local id="3.6"
  local desc="确保/etc/docker的目录权限为755或更多限制性 (Automated)"
  local remediation="You should run the following command: chmod 755 /etc/docker. This sets the permissions for the directory to 755."
  local remediationImpact="None."
  local check="$id - $desc"
  starttestjson "$id" "$desc"

  directory="/etc/docker"
  if [ -d "$directory" ]; then
    if [ "$(stat -c %a $directory)" -le 755 ]; then
      pass -s "$check"
      logcheckresult "PASS"
      return
    fi
    warn -s "$check"
    warn "     * Wrong permissions for $directory"
    logcheckresult "WARN" "Wrong permissions for $directory"
    return
  fi
  info -c "$check"
  info "     * Directory not found"
  logcheckresult "INFO" "Directory not found"
}

check_3_7() {
  local id="3.7"
  local desc="确保镜像仓库证书文件所有权为root:root (Automated)"
  local remediation="You should run the following command: chown root:root /etc/docker/certs.d/<registry-name>/*. This would set the individual ownership and group ownership for the registry certificate files to root."
  local remediationImpact="None."
  local check="$id - $desc"
  starttestjson "$id" "$desc"

  directory="/etc/docker/certs.d/"
  if [ -d "$directory" ]; then
    fail=0
    owners=$(find "$directory" -type f -name '*.crt')
    for p in $owners; do
      if [ "$(stat -c %u "$p")" -ne 0 ]; then
        fail=1
      fi
    done
    if [ $fail -eq 1 ]; then
      warn -s "$check"
      warn "     * Wrong ownership for $directory"
      logcheckresult "WARN" "Wrong ownership for $directory"
      return
    fi
    pass -s "$check"
    logcheckresult "PASS"
    return
  fi
  info -c "$check"
  info "     * Directory not found"
  logcheckresult "INFO" "Directory not found"
}

check_3_8() {
  local id="3.8"
  local desc="确保镜像仓库证书文件权限设置为444或更多限制性 (Automated)"
  local remediation="You should run the following command: chmod 444 /etc/docker/certs.d/<registry-name>/*. This would set the permissions for the registry certificate files to 444."
  local remediationImpact="None."
  local check="$id - $desc"
  starttestjson "$id" "$desc"

  directory="/etc/docker/certs.d/"
  if [ -d "$directory" ]; then
    fail=0
    perms=$(find "$directory" -type f -name '*.crt')
    for p in $perms; do
      if [ "$(stat -c %a "$p")" -gt 444 ]; then
        fail=1
      fi
    done
    if [ $fail -eq 1 ]; then
      warn -s "$check"
      warn "     * Wrong permissions for $directory"
      logcheckresult "WARN" "Wrong permissions for $directory"
      return
    fi
    pass -s "$check"
    logcheckresult "PASS"
    return
  fi
  info -c "$check"
  info "     * Directory not found"
  logcheckresult "INFO" "Directory not found"
}

check_3_9() {
  local id="3.9"
  local desc="确保TLS CA证书文件的所有权为root:root (Automated)"
  local remediation="You should run the following command: chown root:root <path to TLS CA certificate file>. This sets the individual ownership and group ownership for the TLS CA certificate file to root."
  local remediationImpact="None."
  local check="$id - $desc"
  starttestjson "$id" "$desc"

  tlscacert=$(get_docker_effective_command_line_args '--tlscacert' | sed -n 's/.*tlscacert=\([^s]\)/\1/p' | sed 's/--/ --/g' | cut -d " " -f 1)
  if [ -n "$(get_docker_configuration_file_args 'tlscacert')" ]; then
    tlscacert=$(get_docker_configuration_file_args 'tlscacert')
  fi
  if [ -f "$tlscacert" ]; then
    if [ "$(stat -c %u%g "$tlscacert")" -eq 00 ]; then
      pass -s "$check"
      logcheckresult "PASS"
      return
    fi
    warn -s "$check"
    warn "     * Wrong ownership for $tlscacert"
    logcheckresult "WARN" "Wrong ownership for $tlscacert"
    return
  fi
  info -c "$check"
  info "     * No TLS CA certificate found"
  logcheckresult "INFO" "No TLS CA certificate found"
}

check_3_10() {
  local id="3.10"
  local desc="确保TLS CA证书文件权限为444或更多限制性 (Automated)"
  local remediation="You should run the following command: chmod 444 <path to TLS CA certificate file>. This sets the file permissions on the TLS CA file to 444."
  local remediationImpact="None."
  local check="$id - $desc"
  starttestjson "$id" "$desc"

  tlscacert=$(get_docker_effective_command_line_args '--tlscacert' | sed -n 's/.*tlscacert=\([^s]\)/\1/p' | sed 's/--/ --/g' | cut -d " " -f 1)
  if [ -n "$(get_docker_configuration_file_args 'tlscacert')" ]; then
    tlscacert=$(get_docker_configuration_file_args 'tlscacert')
  fi
  if [ -f "$tlscacert" ]; then
    if [ "$(stat -c %a "$tlscacert")" -le 444 ]; then
      pass -s "$check"
      logcheckresult "PASS"
      return
    fi
    warn -s "$check"
    warn "      * Wrong permissions for $tlscacert"
    logcheckresult "WARN" "Wrong permissions for $tlscacert"
    return
  fi
  info -c "$check"
  info "      * No TLS CA certificate found"
  logcheckresult "INFO" "No TLS CA certificate found"
}

check_3_11() {
  local id="3.11"
  local desc="确保Docker服务器证书文件所有权为root:root (Automated)"
  local remediation="You should run the following command: chown root:root <path to Docker server certificate file>. This sets the individual ownership and the group ownership for the Docker server certificate file to root."
  local remediationImpact="None."
  local check="$id - $desc"
  starttestjson "$id" "$desc"

  tlscert=$(get_docker_effective_command_line_args '--tlscert' | sed -n 's/.*tlscert=\([^s]\)/\1/p' | sed 's/--/ --/g' | cut -d " " -f 1)
  if [ -n "$(get_docker_configuration_file_args 'tlscert')" ]; then
    tlscert=$(get_docker_configuration_file_args 'tlscert')
  fi
  if [ -f "$tlscert" ]; then
    if [ "$(stat -c %u%g "$tlscert")" -eq 00 ]; then
      pass -s "$check"
      logcheckresult "PASS"
      return
    fi
    warn -s "$check"
    warn "      * Wrong ownership for $tlscert"
    logcheckresult "WARN" "Wrong ownership for $tlscert"
    return
  fi
  info -c "$check"
  info "      * No TLS Server certificate found"
  logcheckresult "INFO" "No TLS Server certificate found"
}

check_3_12() {
  local id="3.12"
  local desc="确保Docker服务器证书文件权限设置为444或更多限制性 (Automated)"
  local remediation="You should run the following command: chmod 444 <path to Docker server certificate file>. This sets the file permissions of the Docker server certificate file to 444."
  local remediationImpact="None."
  local check="$id - $desc"
  starttestjson "$id" "$desc"

  tlscert=$(get_docker_effective_command_line_args '--tlscert' | sed -n 's/.*tlscert=\([^s]\)/\1/p' | sed 's/--/ --/g' | cut -d " " -f 1)
  if [ -n "$(get_docker_configuration_file_args 'tlscert')" ]; then
    tlscert=$(get_docker_configuration_file_args 'tlscert')
  fi
  if [ -f "$tlscert" ]; then
    if [ "$(stat -c %a "$tlscert")" -le 444 ]; then
      pass -s "$check"
      logcheckresult "PASS"
      return
    fi
    warn -s "$check"
    warn "      * Wrong permissions for $tlscert"
    logcheckresult "WARN" "Wrong permissions for $tlscert"
    return
  fi
  info -c "$check"
  info "      * No TLS Server certificate found"
  logcheckresult "INFO" "No TLS Server certificate found"
}

check_3_13() {
  local id="3.13"
  local desc="确保Docker服务器证书密钥文件的所有权为root:root (Automated)"
  local remediation="You should run the following command: chown root:root <path to Docker server certificate key file>. This sets the individual ownership and group ownership for the Docker server certificate key file to root."
  local remediationImpact="None."
  local check="$id - $desc"
  starttestjson "$id" "$desc"

  tlskey=$(get_docker_effective_command_line_args '--tlskey' | sed -n 's/.*tlskey=\([^s]\)/\1/p' | sed 's/--/ --/g' | cut -d " " -f 1)
  if [ -n "$(get_docker_configuration_file_args 'tlskey')" ]; then
    tlskey=$(get_docker_configuration_file_args 'tlskey')
  fi
  if [ -f "$tlskey" ]; then
    if [ "$(stat -c %u%g "$tlskey")" -eq 00 ]; then
      pass -s "$check"
      logcheckresult "PASS"
      return
    fi
    warn -s "$check"
    warn "      * Wrong ownership for $tlskey"
    logcheckresult "WARN" "Wrong ownership for $tlskey"
    return
  fi
  info -c "$check"
  info "      * No TLS Key found"
  logcheckresult "INFO" "No TLS Key found"
}

check_3_14() {
  local id="3.14"
  local desc="确保Docker服务器证书密钥文件权限为400 (Automated)"
  local remediation="You should run the following command: chmod 400 <path to Docker server certificate key file>. This sets the Docker server certificate key file permissions to 400."
  local remediationImpact="None."
  local check="$id - $desc"
  starttestjson "$id" "$desc"

  tlskey=$(get_docker_effective_command_line_args '--tlskey' | sed -n 's/.*tlskey=\([^s]\)/\1/p' | sed 's/--/ --/g' | cut -d " " -f 1)
  if [ -n "$(get_docker_configuration_file_args 'tlskey')" ]; then
    tlskey=$(get_docker_configuration_file_args 'tlskey')
  fi
  if [ -f "$tlskey" ]; then
    if [ "$(stat -c %a "$tlskey")" -eq 400 ]; then
      pass -s "$check"
      logcheckresult "PASS"
      return
    fi
    warn -s "$check"
    warn "      * Wrong permissions for $tlskey"
    logcheckresult "WARN" "Wrong permissions for $tlskey"
    return
  fi
  info -c "$check"
  info "      * No TLS Key found"
  logcheckresult "INFO" "No TLS Key found"
}

check_3_15() {
  local id="3.15"
  local desc="确保Docker.sock的文件所有权为root:docker (Automated)"
  local remediation="You should run the following command: chown root:docker /var/run/docker.sock. This sets the ownership to root and group ownership to docker for the default Docker socket file."
  local remediationImpact="None."
  local check="$id - $desc"
  starttestjson "$id" "$desc"

  file="/var/run/docker.sock"
  if [ -S "$file" ]; then
    if [ "$(stat -c %U:%G $file)" = 'root:docker' ]; then
      pass -s "$check"
      logcheckresult "PASS"
      return
    fi
    warn -s "$check"
    warn "      * Wrong ownership for $file"
    logcheckresult "WARN" "Wrong ownership for $file"
    return
  fi
  info -c "$check"
  info "      * File not found"
  logcheckresult "INFO" "File not found"
}

check_3_16() {
  local id="3.16"
  local desc="确保Docker.sock的文件权限设置为660或更多限制性 (Automated)"
  local remediation="You should run the following command: chmod 660 /var/run/docker.sock. This sets the file permissions of the Docker socket file to 660."
  local remediationImpact="None."
  local check="$id - $desc"
  starttestjson "$id" "$desc"

  file="/var/run/docker.sock"
  if [ -S "$file" ]; then
    if [ "$(stat -c %a $file)" -le 660 ]; then
      pass -s "$check"
      logcheckresult "PASS"
      return
    fi
    warn -s "$check"
    warn "      * Wrong permissions for $file"
    logcheckresult "WARN" "Wrong permissions for $file"
    return
  fi
  info -c "$check"
  info "      * File not found"
  logcheckresult "INFO" "File not found"
}

check_3_17() {
  local id="3.17"
  local desc="确保daemon.json文件所有权为root:root (Automated)"
  local remediation="You should run the following command: chown root:root /etc/docker/daemon.json. This sets the ownership and group ownership for the file to root."
  local remediationImpact="None."
  local check="$id - $desc"
  starttestjson "$id" "$desc"

  file="/etc/docker/daemon.json"
  if [ -f "$file" ]; then
    if [ "$(stat -c %U:%G $file)" = 'root:root' ]; then
      pass -s "$check"
      logcheckresult "PASS"
      return
    fi
    warn -s "$check"
    warn "      * Wrong ownership for $file"
    logcheckresult "WARN" "Wrong ownership for $file"
    return
  fi
  info -c "$check"
  info "      * File not found"
  logcheckresult "INFO" "File not found"
}

check_3_18() {
  local id="3.18"
  local desc="确保daemon.json文件权限为644或更多限制性 (Automated)"
  local remediation="You should run the following command: chmod 644 /etc/docker/daemon.json. This sets the file permissions for this file to 644."
  local remediationImpact="None."
  local check="$id - $desc"
  starttestjson "$id" "$desc"

  file="/etc/docker/daemon.json"
  if [ -f "$file" ]; then
    if [ "$(stat -c %a $file)" -le 644 ]; then
      pass -s "$check"
      logcheckresult "PASS"
      return
    fi
    warn -s "$check"
    warn "      * Wrong permissions for $file"
    logcheckresult "WARN" "Wrong permissions for $file"
    return
  fi
  info -c "$check"
  info "      * File not found"
  logcheckresult "INFO" "File not found"
}

check_3_19() {
  local id="3.19"
  local desc="确保/etc/default/docker文件的所有权为root:root (Automated)"
  local remediation="You should run the following command: chown root:root /etc/default/docker. This sets the ownership and group ownership of the file to root."
  local remediationImpact="None."
  local check="$id - $desc"
  starttestjson "$id" "$desc"

  file="/etc/default/docker"
  if [ -f "$file" ]; then
    if [ "$(stat -c %U:%G $file)" = 'root:root' ]; then
      pass -s "$check"
      logcheckresult "PASS"
      return
    fi
    warn -s "$check"
    warn "      * Wrong ownership for $file"
    logcheckresult "WARN" "Wrong ownership for $file"
    return
  fi
  info -c "$check"
  info "      * File not found"
  logcheckresult "INFO" "File not found"
}

check_3_20() {
  local id="3.20"
  local desc="确保/etc/sysconfig/docker文件的所有权为root:root (Automated)"
  local remediation="You should run the following command: chmod 644 /etc/sysconfig/docker. This sets the file permissions for this file to 644."
  local remediationImpact="None."
  local check="$id - $desc"
  starttestjson "$id" "$desc"

  file="/etc/sysconfig/docker"
  if [ -f "$file" ]; then
    if [ "$(stat -c %a $file)" -le 644 ]; then
      pass -s "$check"
      logcheckresult "PASS"
      return
    fi
    warn -s "$check"
    warn "      * Wrong permissions for $file"
    logcheckresult "WARN" "Wrong permissions for $file"
    return
  fi
  info -c "$check"
  info "      * File not found"
  logcheckresult "INFO" "File not found"
}

check_3_21() {
  local id="3.21"
  local desc="确保/etc/sysconfig/docker文件权限为644或更多限制性 (Automated)"
  local remediation="You should run the following command: chown root:root /etc/sysconfig/docker. This sets the ownership and group ownership for the file to root."
  local remediationImpact="None."
  local check="$id - $desc"
  starttestjson "$id" "$desc"

  file="/etc/sysconfig/docker"
  if [ -f "$file" ]; then
    if [ "$(stat -c %U:%G $file)" = 'root:root' ]; then
      pass -s "$check"
      logcheckresult "PASS"
      return
    fi
    warn -s "$check"
    warn "      * Wrong ownership for $file"
    logcheckresult "WARN" "Wrong ownership for $file"
    return
  fi
  info -c "$check"
  info "      * File not found"
  logcheckresult "INFO" "File not found"
}

check_3_22() {
  local id="3.22"
  local desc="确保/etc/default/docker文件权限为644或更多限制性 (Automated)"
  local remediation="You should run the following command: chmod 644 /etc/default/docker. This sets the file permissions for this file to 644."
  local remediationImpact="None."
  local check="$id - $desc"
  starttestjson "$id" "$desc"

  file="/etc/default/docker"
  if [ -f "$file" ]; then
    if [ "$(stat -c %a $file)" -le 644 ]; then
      pass -s "$check"
      logcheckresult "PASS"
      return
    fi
    warn -s "$check"
    warn "      * Wrong permissions for $file"
    logcheckresult "WARN" "Wrong permissions for $file"
    return
  fi
  info -c "$check"
  info "      * File not found"
  logcheckresult "INFO" "File not found"
}

check_3_23() {
  local id="3.23"
  local desc="确保containerd socket文件所有权为root:root (Automated)"
  local remediation="You should run the following command: chown root:root /run/containerd/containerd.sock. This sets the ownership and group ownership for the file to root."
  local remediationImpact="None."
  local check="$id - $desc"
  starttestjson "$id" "$desc"

  file="/run/containerd/containerd.sock"
  if [ -S "$file" ]; then
    if [ "$(stat -c %U:%G $file)" = 'root:root' ]; then
      pass -s "$check"
      logcheckresult "PASS"
      return
    fi
    warn -s "$check"
    warn "      * Wrong ownership for $file"
    logcheckresult "WARN" "Wrong ownership for $file"
    return
  fi
  info -c "$check"
  info "      * File not found"
  logcheckresult "INFO" "File not found"
}

check_3_24() {
  local id="3.24"
  local desc="确保Containerd socket文件权限被设置为660或更多限制性 (Automated)"
  local remediation="You should run the following command: chmod 660 /run/containerd/containerd.sock. This sets the file permissions for this file to 660."
  local remediationImpact="None."
  local check="$id - $desc"
  starttestjson "$id" "$desc"

  file="/run/containerd/containerd.sock"
  if [ -S "$file" ]; then
    if [ "$(stat -c %a $file)" -le 660 ]; then
      pass -s "$check"
      logcheckresult "PASS"
      return
    fi
    warn -s "$check"
    warn "      * Wrong permissions for $file"
    logcheckresult "WARN" "Wrong permissions for $file"
    return
  fi
  info -c "$check"
  info "      * File not found"
  logcheckresult "INFO" "File not found"
}

check_3_end() {
  endsectionjson
}
