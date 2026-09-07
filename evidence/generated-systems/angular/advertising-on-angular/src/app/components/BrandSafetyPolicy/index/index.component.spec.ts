
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexBrandSafetyPolicyComponent } from './index.component';
import { BrandSafetyPolicyService } from '../../../services/BrandSafetyPolicy.service';

describe('IndexBrandSafetyPolicyComponent', () => {
  let component: IndexBrandSafetyPolicyComponent;
  let fixture: ComponentFixture<IndexBrandSafetyPolicyComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexBrandSafetyPolicyComponent
      ],
      providers: [
        BrandSafetyPolicyService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexBrandSafetyPolicyComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});