import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { Model_Service } from '../../../services/Model_.service';
import { SubBaseComponent } from '../../Model_/sub.base.component';


@Component({
    selector: 'app-edit-model_',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditModel_Component extends SubBaseComponent implements OnInit {

    title = 'Edit Model_';

    model_Form: FormGroup;
    model_: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: Model_Service,
        private fb: FormBuilder
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

    
    updateModel_(name, taskDescription, Workspace, Versions, FeatureSets, Experiments, Tags, ModelType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateModel_(name, taskDescription, Workspace, Versions, FeatureSets, Experiments, Tags, ModelType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexModel_']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getModel_(params['id']).subscribe(res => {
                this.model_ = res;
            });
        });
    }
}