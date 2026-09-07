import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { PredictionService } from '../../../services/Prediction.service';
import { SubBaseComponent } from '../../Prediction/sub.base.component';


@Component({
    selector: 'app-edit-prediction',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditPredictionComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Prediction';

    predictionForm: FormGroup;
    prediction: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: PredictionService,
        private fb: FormBuilder
) {
        super(http);
        this.predictionForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  referenceKey: ['', Validators.required],
      predictedAt: ['', Validators.required],
      score: ['', Validators.required],
      Endpoint: ['', ],
      ModelVersion: ['', ],
      Dataset: ['', ]
        });
    }

    
    updatePrediction(referenceKey, predictedAt, score, Endpoint, ModelVersion, Dataset): void {
        this.route.params.subscribe((params) => {

                        this.service.updatePrediction(referenceKey, predictedAt, score, Endpoint, ModelVersion, Dataset, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexPrediction']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getPrediction(params['id']).subscribe(res => {
                this.prediction = res;
            });
        });
    }
}