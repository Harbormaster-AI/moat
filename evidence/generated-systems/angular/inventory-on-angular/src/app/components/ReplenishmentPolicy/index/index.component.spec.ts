
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexReplenishmentPolicyComponent } from './index.component';
import { ReplenishmentPolicyService } from '../../../services/ReplenishmentPolicy.service';

describe('IndexReplenishmentPolicyComponent', () => {
  let component: IndexReplenishmentPolicyComponent;
  let fixture: ComponentFixture<IndexReplenishmentPolicyComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexReplenishmentPolicyComponent
      ],
      providers: [
        ReplenishmentPolicyService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexReplenishmentPolicyComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});