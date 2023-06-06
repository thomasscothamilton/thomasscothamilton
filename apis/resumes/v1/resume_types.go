/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type Education struct {
	Institution string   `json:"institution"`
	Area        string   `json:"area"`
	StudyType   string   `json:"studyType"`
	StartDate   string   `json:"startDate"`
	EndDate     string   `json:"endDate"`
	Gpa         string   `json:"gpa"`
	Courses     []string `json:"courses"`
}

type Experience struct {
	Company    string   `json:"company"`
	Position   string   `json:"position"`
	Website    string   `json:"website"`
	StartDate  string   `json:"startDate"`
	EndDate    string   `json:"endDate"`
	Summary    string   `json:"summary"`
	Highlights []string `json:"highlights"`
}

type Skill struct {
	Name     string   `json:"name"`
	Level    string   `json:"level"`
	Keywords []string `json:"keywords"`
}

type Project struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Highlights  []string `json:"highlights"`
	Keywords    []string `json:"keywords"`
	StartDate   string   `json:"startDate"`
	EndDate     string   `json:"endDate"`
	Url         string   `json:"url"`
}

type Award struct {
	Title   string `json:"title"`
	Date    string `json:"date"`
	Awarder string `json:"awarder"`
	Summary string `json:"summary"`
}

type Publication struct {
	Name        string `json:"name"`
	Publisher   string `json:"publisher"`
	ReleaseDate string `json:"releaseDate"`
	Website     string `json:"website"`
	Summary     string `json:"summary"`
}

type Language struct {
	Language string `json:"language"`
	Fluency  string `json:"fluency"`
}

type Interest struct {
	Name     string   `json:"name"`
	Keywords []string `json:"keywords"`
}

type Reference struct {
	Name      string `json:"name"`
	Reference string `json:"reference"`
}

// ResumeSpec defines the desired state of Resume
type ResumeSpec struct {
	Name         string        `json:"name"`
	Email        string        `json:"email"`
	Phone        string        `json:"phone"`
	Website      string        `json:"website"`
	GitHub       string        `json:"github"`
	LinkedIn     string        `json:"linkedin"`
	Summary      string        `json:"summary"`
	Education    []Education   `json:"education"`
	Experience   []Experience  `json:"experience"`
	Skills       []Skill       `json:"skills"`
	Projects     []Project     `json:"projects"`
	Awards       []Award       `json:"awards"`
	Publications []Publication `json:"publications"`
	Languages    []Language    `json:"languages"`
	Interests    []Interest    `json:"interests"`
	References   []Reference   `json:"references"`
}

// ResumeStatus defines the observed state of Resume
type ResumeStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Resume is the Schema for the resumes API
type Resume struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ResumeSpec   `json:"spec,omitempty"`
	Status ResumeStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ResumeList contains a list of Resume
type ResumeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Resume `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Resume{}, &ResumeList{})
}
