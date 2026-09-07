import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { LabResultService } from '../../../services/LabResult.service';
import { LabResult } from '../../../models/LabResult';
import { SubBaseComponent } from '../../LabResult/sub.base.component';

@Component({
    selector: 'app-create-labResult',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateLabResultComponent extends SubBaseComponent implements OnInit {

    title = 'Add LabResult';

    labResultForm: FormGroup;
    labResult: LabResult;

    constructor( http: HttpClient,
        private labResultService: LabResultService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.labResultForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  resultCode: ['', Validators.required],
      issuedDate: ['', Validators.required],
      LaboratoryOrder: ['', ],
      Observations: ['', ],
      Laboratory: ['', ],
      Status: ['', ]
        });
    }

    
    addLabResult(resultCode, issuedDate, LaboratoryOrder, Observations, Laboratory, Status): void {
        this.labResultService
        .addLabResult(resultCode, issuedDate, LaboratoryOrder, Observations, Laboratory, Status)
            .subscribe(() => {
                this.router.navigate(['/indexLabResult']);
            });
    }

    ngOnInit(): void {
    }
}