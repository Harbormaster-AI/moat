
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexBillingProfileComponent } from './index.component';
import { BillingProfileService } from '../../../services/BillingProfile.service';

describe('IndexBillingProfileComponent', () => {
  let component: IndexBillingProfileComponent;
  let fixture: ComponentFixture<IndexBillingProfileComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexBillingProfileComponent
      ],
      providers: [
        BillingProfileService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexBillingProfileComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});