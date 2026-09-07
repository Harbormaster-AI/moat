import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { EquityGrantService } from '../../../services/EquityGrant.service';
import { SubBaseComponent } from '../../EquityGrant/sub.base.component';


@Component({
    selector: 'app-edit-equityGrant',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditEquityGrantComponent extends SubBaseComponent implements OnInit {

    title = 'Edit EquityGrant';

    equityGrantForm: FormGroup;
    equityGrant: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: EquityGrantService,
        private fb: FormBuilder
) {
        super(http);
        this.equityGrantForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  grantId: ['', Validators.required],
      grantedUnits: ['', Validators.required],
      vestingStart: ['', Validators.required],
      CompensationPackage: ['', ],
      GrantType: ['', ]
        });
    }

    
    updateEquityGrant(grantId, grantedUnits, vestingStart, CompensationPackage, GrantType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateEquityGrant(grantId, grantedUnits, vestingStart, CompensationPackage, GrantType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexEquityGrant']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getEquityGrant(params['id']).subscribe(res => {
                this.equityGrant = res;
            });
        });
    }
}