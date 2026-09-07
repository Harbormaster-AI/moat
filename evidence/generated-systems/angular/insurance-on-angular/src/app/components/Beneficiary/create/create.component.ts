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
      share: ['', Validators.required],
      Policy: ['', ],
      Customer: ['', ],
      Relationship: ['', ]
        });
    }

    
    addBeneficiary(name, share, Policy, Customer, Relationship): void {
        this.beneficiaryService
        .addBeneficiary(name, share, Policy, Customer, Relationship)
            .subscribe(() => {
                this.router.navigate(['/indexBeneficiary']);
            });
    }

    ngOnInit(): void {
    }
}