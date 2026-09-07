
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EditBrandSafetyPolicyComponent } from './edit.component';
import { BrandSafetyPolicyService } from '../../../services/BrandSafetyPolicy.service';

describe('EditBrandSafetyPolicyComponent', () => {
  let component: EditBrandSafetyPolicyComponent;
  let fixture: ComponentFixture<EditBrandSafetyPolicyComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        EditBrandSafetyPolicyComponent
      ],
      providers: [
        BrandSafetyPolicyService,
        {
          provide: ActivatedRoute,
          useValue: {
            params: of({ id: '1' })
          }
        },
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(EditBrandSafetyPolicyComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});