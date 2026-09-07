import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { TaxWithholdingService } from '../../../services/TaxWithholding.service';
import { TaxWithholding } from '../../../models/TaxWithholding';
import { SubBaseComponent } from '../../TaxWithholding/sub.base.component';

@Component({
    selector: 'app-create-taxWithholding',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateTaxWithholdingComponent extends SubBaseComponent implements OnInit {

    title = 'Add TaxWithholding';

    taxWithholdingForm: FormGroup;
    taxWithholding: TaxWithholding;

    constructor( http: HttpClient,
        private taxWithholdingService: TaxWithholdingService,
        private fb: FormBuilder,
        private router: Router
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

    
    addTaxWithholding(taxId, allowances, additionalAmount, Employee, FilingStatus): void {
        this.taxWithholdingService
        .addTaxWithholding(taxId, allowances, additionalAmount, Employee, FilingStatus)
            .subscribe(() => {
                this.router.navigate(['/indexTaxWithholding']);
            });
    }

    ngOnInit(): void {
    }
}