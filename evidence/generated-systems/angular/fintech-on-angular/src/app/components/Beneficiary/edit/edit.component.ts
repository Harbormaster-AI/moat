import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { BeneficiaryService } from '../../../services/Beneficiary.service';
import { SubBaseComponent } from '../../Beneficiary/sub.base.component';


@Component({
    selector: 'app-edit-beneficiary',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditBeneficiaryComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Beneficiary';

    beneficiaryForm: FormGroup;
    beneficiary: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: BeneficiaryService,
        private fb: FormBuilder
) {
        super(http);
        this.beneficiaryForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      accountIdentifier: ['', Validators.required],
      iban: ['', Validators.required],
      bic: ['', Validators.required],
      address: ['', Validators.required],
      Customer: ['', ]
        });
    }

    
    updateBeneficiary(name, accountIdentifier, iban, bic, address, Customer): void {
        this.route.params.subscribe((params) => {

                        this.service.updateBeneficiary(name, accountIdentifier, iban, bic, address, Customer, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexBeneficiary']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getBeneficiary(params['id']).subscribe(res => {
                this.beneficiary = res;
            });
        });
    }
}