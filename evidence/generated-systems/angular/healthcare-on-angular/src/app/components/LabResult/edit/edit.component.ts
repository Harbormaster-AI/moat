import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { LabResultService } from '../../../services/LabResult.service';
import { SubBaseComponent } from '../../LabResult/sub.base.component';


@Component({
    selector: 'app-edit-labResult',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditLabResultComponent extends SubBaseComponent implements OnInit {

    title = 'Edit LabResult';

    labResultForm: FormGroup;
    labResult: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: LabResultService,
        private fb: FormBuilder
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

    
    updateLabResult(resultCode, issuedDate, LaboratoryOrder, Observations, Laboratory, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateLabResult(resultCode, issuedDate, LaboratoryOrder, Observations, Laboratory, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexLabResult']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getLabResult(params['id']).subscribe(res => {
                this.labResult = res;
            });
        });
    }
}