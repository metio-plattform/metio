// SPDX-License-Identifier: Apache-2.0

//go:generate controller-gen object paths=.
//go:generate controller-gen crd paths=. output:crd:artifacts:config=../../../../config/crd/bases

package v1alpha1
