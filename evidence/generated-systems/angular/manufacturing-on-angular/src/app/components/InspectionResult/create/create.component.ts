import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { InspectionResultService } from '../../../services/InspectionResult.service';
import { InspectionResult } from '../../../models/InspectionResult';
import { SubBaseComponent } from '../../InspectionResult/sub.base.component';

@Component({
    selector: 'app-create-inspectionResult',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateInspectionResultComponent extends SubBaseComponent implements OnInit {

    title = 'Add InspectionResult';

    inspectionResultForm: FormGroup;
    inspectionResult: InspectionResult;

    constructor( http: HttpClient,
        private inspectionResultService: InspectionResultService,
        private fb: FormBuilder,
        private router: Router
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

    
    addInspectionResult(resultValue, recordedOn, notes, InspectionLot, Characteristic, ResultStatus): void {
        this.inspectionResultService
        .addInspectionResult(resultValue, recordedOn, notes, InspectionLot, Characteristic, ResultStatus)
            .subscribe(() => {
                this.router.navigate(['/indexInspectionResult']);
            });
    }

    ngOnInit(): void {
    }
}