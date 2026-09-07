import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { FeatureService } from '../../../services/Feature.service';
import { SubBaseComponent } from '../../Feature/sub.base.component';


@Component({
    selector: 'app-edit-feature',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditFeatureComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Feature';

    featureForm: FormGroup;
    feature: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: FeatureService,
        private fb: FormBuilder
) {
        super(http);
        this.featureForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      description: ['', Validators.required],
      FeatureSet: ['', ],
      SourceDatasets: ['', ],
      Models: ['', ],
      TrainingRuns: ['', ],
      DataType: ['', ]
        });
    }

    
    updateFeature(name, description, FeatureSet, SourceDatasets, Models, TrainingRuns, DataType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateFeature(name, description, FeatureSet, SourceDatasets, Models, TrainingRuns, DataType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexFeature']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getFeature(params['id']).subscribe(res => {
                this.feature = res;
            });
        });
    }
}