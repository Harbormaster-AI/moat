import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { ThirdPartyService } from '../../../services/ThirdParty.service';
import { SubBaseComponent } from '../../ThirdParty/sub.base.component';


@Component({
    selector: 'app-edit-thirdParty',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditThirdPartyComponent extends SubBaseComponent implements OnInit {

    title = 'Edit ThirdParty';

    thirdPartyForm: FormGroup;
    thirdParty: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: ThirdPartyService,
        private fb: FormBuilder
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

    
    updateThirdParty(name, taxId, address, Subrogations, PartyType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateThirdParty(name, taxId, address, Subrogations, PartyType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexThirdParty']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getThirdParty(params['id']).subscribe(res => {
                this.thirdParty = res;
            });
        });
    }
}