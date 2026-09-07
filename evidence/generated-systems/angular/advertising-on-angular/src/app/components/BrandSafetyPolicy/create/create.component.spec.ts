
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateBrandSafetyPolicyComponent } from './create.component';
import { BrandSafetyPolicyService } from '../../../services/BrandSafetyPolicy.service';
import { Router } from '@angular/router';

describe('CreateBrandSafetyPolicyComponent', () => {
  let component: CreateBrandSafetyPolicyComponent;
  let fixture: ComponentFixture<CreateBrandSafetyPolicyComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateBrandSafetyPolicyComponent
      ],
      providers: [
        BrandSafetyPolicyService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateBrandSafetyPolicyComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});