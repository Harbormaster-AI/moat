import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { CardTokenizationService } from '../../../services/CardTokenization.service';
import { SubBaseComponent } from '../../CardTokenization/sub.base.component';


@Component({
    selector: 'app-edit-cardTokenization',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditCardTokenizationComponent extends SubBaseComponent implements OnInit {

    title = 'Edit CardTokenization';

    cardTokenizationForm: FormGroup;
    cardTokenization: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: CardTokenizationService,
        private fb: FormBuilder
) {
        super(http);
        this.cardTokenizationForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  tokenReference: ['', Validators.required],
      createdAt: ['', Validators.required],
      Card: ['', ],
      WalletProvider: ['', ],
      Status: ['', ]
        });
    }

    
    updateCardTokenization(tokenReference, createdAt, Card, WalletProvider, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateCardTokenization(tokenReference, createdAt, Card, WalletProvider, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexCardTokenization']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getCardTokenization(params['id']).subscribe(res => {
                this.cardTokenization = res;
            });
        });
    }
}