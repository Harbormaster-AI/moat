import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { TaxWithholdingService } from '../../../services/TaxWithholding.service';
import { SubBaseComponent } from '../../TaxWithholding/sub.base.component';


@Component({
    selector: 'app-edit-taxWithholding',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditTaxWithholdingComponent extends SubBaseComponent implements OnInit {

    title = 'Edit TaxWithholding';

    taxWithholdingForm: FormGroup;
    taxWithholding: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: TaxWithholdingService,
        private fb: FormBuilder
) {
        super(http);
        this.taxWithholdingForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  taxId: ['', Validators.required],
      allowances: ['', Validators.required],
      additionalAmount: ['', Validators.required],
      Employee: ['', ],
      FilingStatus: ['', ]
        });
    }

    
    updateTaxWithholding(taxId, allowances, additionalAmount, Employee, FilingStatus): void {
        this.route.params.subscribe((params) => {

                        this.service.updateTaxWithholding(taxId, allowances, additionalAmount, Employee, FilingStatus, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexTaxWithholding']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getTaxWithholding(params['id']).subscribe(res => {
                this.taxWithholding = res;
            });
        });
    }
}