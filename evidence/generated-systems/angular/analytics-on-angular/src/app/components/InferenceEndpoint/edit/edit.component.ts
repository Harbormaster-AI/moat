import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { InferenceEndpointService } from '../../../services/InferenceEndpoint.service';
import { SubBaseComponent } from '../../InferenceEndpoint/sub.base.component';


@Component({
    selector: 'app-edit-inferenceEndpoint',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditInferenceEndpointComponent extends SubBaseComponent implements OnInit {

    title = 'Edit InferenceEndpoint';

    inferenceEndpointForm: FormGroup;
    inferenceEndpoint: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: InferenceEndpointService,
        private fb: FormBuilder
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

    
    updateInferenceEndpoint(name, endpointUrl, trafficShare, ModelVersion, Workspace, Predictions, Mode): void {
        this.route.params.subscribe((params) => {

                        this.service.updateInferenceEndpoint(name, endpointUrl, trafficShare, ModelVersion, Workspace, Predictions, Mode, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexInferenceEndpoint']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getInferenceEndpoint(params['id']).subscribe(res => {
                this.inferenceEndpoint = res;
            });
        });
    }
}