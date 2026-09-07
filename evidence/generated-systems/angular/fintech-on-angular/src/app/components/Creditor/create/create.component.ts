import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { CreditorService } from '../../../services/Creditor.service';
import { Creditor } from '../../../models/Creditor';
import { SubBaseComponent } from '../../Creditor/sub.base.component';

@Component({
    selector: 'app-create-creditor',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateCreditorComponent extends SubBaseComponent implements OnInit {

    title = 'Add Creditor';

    creditorForm: FormGroup;
    creditor: Creditor;

    constructor( http: HttpClient,
        private creditorService: CreditorService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.creditorForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      bic: ['', Validators.required],
      address: ['', Validators.required],
      Mandates: ['', ]
        });
    }

    
    addCreditor(name, bic, address, Mandates): void {
        this.creditorService
        .addCreditor(name, bic, address, Mandates)
            .subscribe(() => {
                this.router.navigate(['/indexCreditor']);
            });
    }

    ngOnInit(): void {
    }
}