import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { FeatureSetService } from '../../../services/FeatureSet.service';
import { SubBaseComponent } from '../../FeatureSet/sub.base.component';


@Component({
    selector: 'app-edit-featureSet',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditFeatureSetComponent extends SubBaseComponent implements OnInit {

    title = 'Edit FeatureSet';

    featureSetForm: FormGroup;
    featureSet: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: FeatureSetService,
        private fb: FormBuilder
) {
        super(http);
        this.featureSetForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      refreshSchedule: ['', Validators.required],
      Workspace: ['', ],
      Features: ['', ],
      Datasets: ['', ],
      Models: ['', ],
      ModelVersions: ['', ],
      Tags: ['', ],
      StoreType: ['', ]
        });
    }

    
    updateFeatureSet(name, refreshSchedule, Workspace, Features, Datasets, Models, ModelVersions, Tags, StoreType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateFeatureSet(name, refreshSchedule, Workspace, Features, Datasets, Models, ModelVersions, Tags, StoreType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexFeatureSet']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getFeatureSet(params['id']).subscribe(res => {
                this.featureSet = res;
            });
        });
    }
}