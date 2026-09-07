import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { FeatureService } from '../../../services/Feature.service';
import { Feature } from '../../../models/Feature';
import { SubBaseComponent } from '../../Feature/sub.base.component';

@Component({
    selector: 'app-create-feature',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateFeatureComponent extends SubBaseComponent implements OnInit {

    title = 'Add Feature';

    featureForm: FormGroup;
    feature: Feature;

    constructor( http: HttpClient,
        private featureService: FeatureService,
        private fb: FormBuilder,
        private router: Router
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

    
    addFeature(name, description, FeatureSet, SourceDatasets, Models, TrainingRuns, DataType): void {
        this.featureService
        .addFeature(name, description, FeatureSet, SourceDatasets, Models, TrainingRuns, DataType)
            .subscribe(() => {
                this.router.navigate(['/indexFeature']);
            });
    }

    ngOnInit(): void {
    }
}