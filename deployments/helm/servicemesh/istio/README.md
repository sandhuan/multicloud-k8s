#/*
# * Copyright 2019 Intel Corporation, Inc
# *
# * Licensed under the Apache License, Version 2.0 (the "License");
# * you may not use this file except in compliance with the License.
# * You may obtain a copy of the License at
# *
# *     http://www.apache.org/licenses/LICENSE-2.0
# *
# * Unless required by applicable law or agreed to in writing, software
# * distributed under the License is distributed on an "AS IS" BASIS,
# * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# * See the License for the specific language governing permissions and
# * limitations under the License.
# */

# Steps for Instaling Istio with Istio- Operator

# Step 1 - Add the helm chart to install Istio in sds configuration
# NOTE - Edit the namespaces in istio/istio-instance/values.yaml
# to enable istio-injection
helm install istio-instance --name istio --namespace istio-system
