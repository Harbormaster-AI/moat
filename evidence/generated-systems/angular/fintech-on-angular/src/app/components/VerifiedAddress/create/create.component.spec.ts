
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateVerifiedAddressComponent } from './create.component';
import { VerifiedAddressService } from '../../../services/VerifiedAddress.service';
import { Router } from '@angular/router';

describe('CreateVerifiedAddressComponent', () => {
  let component: CreateVerifiedAddressComponent;
  let fixture: ComponentFixture<CreateVerifiedAddressComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateVerifiedAddressComponent
      ],
      providers: [
        VerifiedAddressService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateVerifiedAddressComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});