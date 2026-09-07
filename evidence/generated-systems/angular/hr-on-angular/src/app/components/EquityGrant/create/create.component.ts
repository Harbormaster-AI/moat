import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { EquityGrantService } from '../../../services/EquityGrant.service';
import { EquityGrant } from '../../../models/EquityGrant';
import { SubBaseComponent } from '../../EquityGrant/sub.base.component';

@Component({
    selector: 'app-create-equityGrant',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateEquityGrantComponent extends SubBaseComponent implements OnInit {

    title = 'Add EquityGrant';

    equityGrantForm: FormGroup;
    equityGrant: EquityGrant;

    constructor( http: HttpClient,
        private equityGrantService: EquityGrantService,
        private fb: FormBuilder,
        private router: Router
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

    
    addEquityGrant(grantId, grantedUnits, vestingStart, CompensationPackage, GrantType): void {
        this.equityGrantService
        .addEquityGrant(grantId, grantedUnits, vestingStart, CompensationPackage, GrantType)
            .subscribe(() => {
                this.router.navigate(['/indexEquityGrant']);
            });
    }

    ngOnInit(): void {
    }
}