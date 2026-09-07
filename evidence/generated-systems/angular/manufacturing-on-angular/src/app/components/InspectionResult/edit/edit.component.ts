import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { InspectionResultService } from '../../../services/InspectionResult.service';
import { SubBaseComponent } from '../../InspectionResult/sub.base.component';


@Component({
    selector: 'app-edit-inspectionResult',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditInspectionResultComponent extends SubBaseComponent implements OnInit {

    title = 'Edit InspectionResult';

    inspectionResultForm: FormGroup;
    inspectionResult: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: InspectionResultService,
        private fb: FormBuilder
) {
        super(http);
        this.inspectionResultForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  resultValue: ['', Validators.required],
      recordedOn: ['', Validators.required],
      notes: ['', Validators.required],
      InspectionLot: ['', ],
      Characteristic: ['', ],
      ResultStatus: ['', ]
        });
    }

    
    updateInspectionResult(resultValue, recordedOn, notes, InspectionLot, Characteristic, ResultStatus): void {
        this.route.params.subscribe((params) => {

                        this.service.updateInspectionResult(resultValue, recordedOn, notes, InspectionLot, Characteristic, ResultStatus, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexInspectionResult']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getInspectionResult(params['id']).subscribe(res => {
                this.inspectionResult = res;
            });
        });
    }
}