
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexBeneficiaryComponent } from './index.component';
import { BeneficiaryService } from '../../../services/Beneficiary.service';

describe('IndexBeneficiaryComponent', () => {
  let component: IndexBeneficiaryComponent;
  let fixture: ComponentFixture<IndexBeneficiaryComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexBeneficiaryComponent
      ],
      providers: [
        BeneficiaryService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexBeneficiaryComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});