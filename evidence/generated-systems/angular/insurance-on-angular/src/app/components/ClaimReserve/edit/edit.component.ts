import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { ClaimReserveService } from '../../../services/ClaimReserve.service';
import { SubBaseComponent } from '../../ClaimReserve/sub.base.component';


@Component({
    selector: 'app-edit-claimReserve',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditClaimReserveComponent extends SubBaseComponent implements OnInit {

    title = 'Edit ClaimReserve';

    claimReserveForm: FormGroup;
    claimReserve: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: ClaimReserveService,
        private fb: FormBuilder
) {
        super(http);
        this.claimReserveForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  amount: ['', Validators.required],
      setDate: ['', Validators.required],
      Claim: ['', ],
      Exposure: ['', ],
      ReserveType: ['', ],
      Status: ['', ]
        });
    }

    
    updateClaimReserve(amount, setDate, Claim, Exposure, ReserveType, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateClaimReserve(amount, setDate, Claim, Exposure, ReserveType, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexClaimReserve']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getClaimReserve(params['id']).subscribe(res => {
                this.claimReserve = res;
            });
        });
    }
}