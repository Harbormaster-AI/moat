import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { FeatureSetService } from '../../../services/FeatureSet.service';
import { FeatureSet } from '../../../models/FeatureSet';
import { SubBaseComponent } from '../../FeatureSet/sub.base.component';

@Component({
    selector: 'app-create-featureSet',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateFeatureSetComponent extends SubBaseComponent implements OnInit {

    title = 'Add FeatureSet';

    featureSetForm: FormGroup;
    featureSet: FeatureSet;

    constructor( http: HttpClient,
        private featureSetService: FeatureSetService,
        private fb: FormBuilder,
        private router: Router
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

    
    addFeatureSet(name, refreshSchedule, Workspace, Features, Datasets, Models, ModelVersions, Tags, StoreType): void {
        this.featureSetService
        .addFeatureSet(name, refreshSchedule, Workspace, Features, Datasets, Models, ModelVersions, Tags, StoreType)
            .subscribe(() => {
                this.router.navigate(['/indexFeatureSet']);
            });
    }

    ngOnInit(): void {
    }
}