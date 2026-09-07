import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Model_Service } from '../../../services/Model_.service';
import { Model_ } from '../../../models/Model_';
import { SubBaseComponent } from '../../Model_/sub.base.component';

@Component({
    selector: 'app-create-model_',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateModel_Component extends SubBaseComponent implements OnInit {

    title = 'Add Model_';

    model_Form: FormGroup;
    model_: Model_;

    constructor( http: HttpClient,
        private model_Service: Model_Service,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.model_Form = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      taskDescription: ['', Validators.required],
      Workspace: ['', ],
      Versions: ['', ],
      FeatureSets: ['', ],
      Experiments: ['', ],
      Tags: ['', ],
      ModelType: ['', ]
        });
    }

    
    addModel_(name, taskDescription, Workspace, Versions, FeatureSets, Experiments, Tags, ModelType): void {
        this.model_Service
        .addModel_(name, taskDescription, Workspace, Versions, FeatureSets, Experiments, Tags, ModelType)
            .subscribe(() => {
                this.router.navigate(['/indexModel_']);
            });
    }

    ngOnInit(): void {
    }
}