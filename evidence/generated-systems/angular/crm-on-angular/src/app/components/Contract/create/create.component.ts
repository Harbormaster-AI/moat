import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ContractService } from '../../../services/Contract.service';
import { Contract } from '../../../models/Contract';
import { SubBaseComponent } from '../../Contract/sub.base.component';

@Component({
    selector: 'app-create-contract',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateContractComponent extends SubBaseComponent implements OnInit {

    title = 'Add Contract';

    contractForm: FormGroup;
    contract: Contract;

    constructor( http: HttpClient,
        private contractService: ContractService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.contractForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  contractNumber: ['', Validators.required],
      startDate: ['', Validators.required],
      endDate: ['', Validators.required],
      renewalTermMonths: ['', Validators.required],
      autoRenew: ['', Validators.required],
      Organization: ['', ],
      Account: ['', ],
      Owner: ['', ],
      Orders: ['', ],
      Cases: ['', ],
      Status: ['', ]
        });
    }

    
    addContract(contractNumber, startDate, endDate, renewalTermMonths, autoRenew, Organization, Account, Owner, Orders, Cases, Status): void {
        this.contractService
        .addContract(contractNumber, startDate, endDate, renewalTermMonths, autoRenew, Organization, Account, Owner, Orders, Cases, Status)
            .subscribe(() => {
                this.router.navigate(['/indexContract']);
            });
    }

    ngOnInit(): void {
    }
}