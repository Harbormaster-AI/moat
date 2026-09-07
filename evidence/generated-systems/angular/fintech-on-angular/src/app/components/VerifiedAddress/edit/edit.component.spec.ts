
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EditVerifiedAddressComponent } from './edit.component';
import { VerifiedAddressService } from '../../../services/VerifiedAddress.service';

describe('EditVerifiedAddressComponent', () => {
  let component: EditVerifiedAddressComponent;
  let fixture: ComponentFixture<EditVerifiedAddressComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        EditVerifiedAddressComponent
      ],
      providers: [
        VerifiedAddressService,
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

    fixture = TestBed.createComponent(EditVerifiedAddressComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});