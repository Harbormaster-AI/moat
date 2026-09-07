
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexBillingAccountComponent } from './index.component';
import { BillingAccountService } from '../../../services/BillingAccount.service';

describe('IndexBillingAccountComponent', () => {
  let component: IndexBillingAccountComponent;
  let fixture: ComponentFixture<IndexBillingAccountComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexBillingAccountComponent
      ],
      providers: [
        BillingAccountService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexBillingAccountComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});