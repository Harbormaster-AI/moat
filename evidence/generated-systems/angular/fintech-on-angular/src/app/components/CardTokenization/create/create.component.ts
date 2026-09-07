import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { CardTokenizationService } from '../../../services/CardTokenization.service';
import { CardTokenization } from '../../../models/CardTokenization';
import { SubBaseComponent } from '../../CardTokenization/sub.base.component';

@Component({
    selector: 'app-create-cardTokenization',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateCardTokenizationComponent extends SubBaseComponent implements OnInit {

    title = 'Add CardTokenization';

    cardTokenizationForm: FormGroup;
    cardTokenization: CardTokenization;

    constructor( http: HttpClient,
        private cardTokenizationService: CardTokenizationService,
        private fb: FormBuilder,
        private router: Router
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

    
    addCardTokenization(tokenReference, createdAt, Card, WalletProvider, Status): void {
        this.cardTokenizationService
        .addCardTokenization(tokenReference, createdAt, Card, WalletProvider, Status)
            .subscribe(() => {
                this.router.navigate(['/indexCardTokenization']);
            });
    }

    ngOnInit(): void {
    }
}