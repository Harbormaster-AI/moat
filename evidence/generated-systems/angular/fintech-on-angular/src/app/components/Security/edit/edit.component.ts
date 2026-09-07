import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { SecurityService } from '../../../services/Security.service';
import { SubBaseComponent } from '../../Security/sub.base.component';


@Component({
    selector: 'app-edit-security',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditSecurityComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Security';

    securityForm: FormGroup;
    security: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: SecurityService,
        private fb: FormBuilder
) {
        super(http);
        this.securityForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  symbol: ['', Validators.required],
      isin: ['', Validators.required],
      cusip: ['', Validators.required],
      currency: ['', Validators.required],
      Positions: ['', ],
      Trades: ['', ],
      Orders: ['', ],
      SecurityType: ['', ]
        });
    }

    
    updateSecurity(symbol, isin, cusip, currency, Positions, Trades, Orders, SecurityType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateSecurity(symbol, isin, cusip, currency, Positions, Trades, Orders, SecurityType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexSecurity']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getSecurity(params['id']).subscribe(res => {
                this.security = res;
            });
        });
    }
}