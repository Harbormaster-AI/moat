import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { SecurityService } from '../../../services/Security.service';
import { Security } from '../../../models/Security';
import { SubBaseComponent } from '../../Security/sub.base.component';

@Component({
    selector: 'app-create-security',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateSecurityComponent extends SubBaseComponent implements OnInit {

    title = 'Add Security';

    securityForm: FormGroup;
    security: Security;

    constructor( http: HttpClient,
        private securityService: SecurityService,
        private fb: FormBuilder,
        private router: Router
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

    
    addSecurity(symbol, isin, cusip, currency, Positions, Trades, Orders, SecurityType): void {
        this.securityService
        .addSecurity(symbol, isin, cusip, currency, Positions, Trades, Orders, SecurityType)
            .subscribe(() => {
                this.router.navigate(['/indexSecurity']);
            });
    }

    ngOnInit(): void {
    }
}