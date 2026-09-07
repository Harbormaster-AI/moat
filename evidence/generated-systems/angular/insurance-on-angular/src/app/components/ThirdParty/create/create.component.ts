import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ThirdPartyService } from '../../../services/ThirdParty.service';
import { ThirdParty } from '../../../models/ThirdParty';
import { SubBaseComponent } from '../../ThirdParty/sub.base.component';

@Component({
    selector: 'app-create-thirdParty',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateThirdPartyComponent extends SubBaseComponent implements OnInit {

    title = 'Add ThirdParty';

    thirdPartyForm: FormGroup;
    thirdParty: ThirdParty;

    constructor( http: HttpClient,
        private thirdPartyService: ThirdPartyService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.thirdPartyForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      taxId: ['', Validators.required],
      address: ['', Validators.required],
      Subrogations: ['', ],
      PartyType: ['', ]
        });
    }

    
    addThirdParty(name, taxId, address, Subrogations, PartyType): void {
        this.thirdPartyService
        .addThirdParty(name, taxId, address, Subrogations, PartyType)
            .subscribe(() => {
                this.router.navigate(['/indexThirdParty']);
            });
    }

    ngOnInit(): void {
    }
}