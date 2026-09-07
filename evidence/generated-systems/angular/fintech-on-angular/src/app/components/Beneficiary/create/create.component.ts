import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { BeneficiaryService } from '../../../services/Beneficiary.service';
import { Beneficiary } from '../../../models/Beneficiary';
import { SubBaseComponent } from '../../Beneficiary/sub.base.component';

@Component({
    selector: 'app-create-beneficiary',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateBeneficiaryComponent extends SubBaseComponent implements OnInit {

    title = 'Add Beneficiary';

    beneficiaryForm: FormGroup;
    beneficiary: Beneficiary;

    constructor( http: HttpClient,
        private beneficiaryService: BeneficiaryService,
        private fb: FormBuilder,
        private router: Router
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

    
    addBeneficiary(name, accountIdentifier, iban, bic, address, Customer): void {
        this.beneficiaryService
        .addBeneficiary(name, accountIdentifier, iban, bic, address, Customer)
            .subscribe(() => {
                this.router.navigate(['/indexBeneficiary']);
            });
    }

    ngOnInit(): void {
    }
}