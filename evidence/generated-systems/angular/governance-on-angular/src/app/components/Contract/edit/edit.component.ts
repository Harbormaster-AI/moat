import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { ContractService } from '../../../services/Contract.service';
import { SubBaseComponent } from '../../Contract/sub.base.component';


@Component({
    selector: 'app-edit-contract',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditContractComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Contract';

    contractForm: FormGroup;
    contract: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: ContractService,
        private fb: FormBuilder
) {
        super(http);
        this.contractForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  title: ['', Validators.required],
      effectiveDate: ['', Validators.required],
      expiryDate: ['', Validators.required],
      repositoryUrl: ['', Validators.required],
      ThirdParty: ['', ],
      Obligations: ['', ],
      DataProcessingActivities: ['', ],
      Matter: ['', ],
      Status: ['', ]
        });
    }

    
    updateContract(title, effectiveDate, expiryDate, repositoryUrl, ThirdParty, Obligations, DataProcessingActivities, Matter, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateContract(title, effectiveDate, expiryDate, repositoryUrl, ThirdParty, Obligations, DataProcessingActivities, Matter, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexContract']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getContract(params['id']).subscribe(res => {
                this.contract = res;
            });
        });
    }
}