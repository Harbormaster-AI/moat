import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { InferenceEndpointService } from '../../../services/InferenceEndpoint.service';
import { InferenceEndpoint } from '../../../models/InferenceEndpoint';
import { SubBaseComponent } from '../../InferenceEndpoint/sub.base.component';

@Component({
    selector: 'app-create-inferenceEndpoint',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateInferenceEndpointComponent extends SubBaseComponent implements OnInit {

    title = 'Add InferenceEndpoint';

    inferenceEndpointForm: FormGroup;
    inferenceEndpoint: InferenceEndpoint;

    constructor( http: HttpClient,
        private inferenceEndpointService: InferenceEndpointService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.inferenceEndpointForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      endpointUrl: ['', Validators.required],
      trafficShare: ['', Validators.required],
      ModelVersion: ['', ],
      Workspace: ['', ],
      Predictions: ['', ],
      Mode: ['', ]
        });
    }

    
    addInferenceEndpoint(name, endpointUrl, trafficShare, ModelVersion, Workspace, Predictions, Mode): void {
        this.inferenceEndpointService
        .addInferenceEndpoint(name, endpointUrl, trafficShare, ModelVersion, Workspace, Predictions, Mode)
            .subscribe(() => {
                this.router.navigate(['/indexInferenceEndpoint']);
            });
    }

    ngOnInit(): void {
    }
}