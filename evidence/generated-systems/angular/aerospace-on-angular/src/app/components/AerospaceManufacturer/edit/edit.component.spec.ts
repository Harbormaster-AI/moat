
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EditAerospaceManufacturerComponent } from './edit.component';
import { AerospaceManufacturerService } from '../../../services/AerospaceManufacturer.service';

describe('EditAerospaceManufacturerComponent', () => {
  let component: EditAerospaceManufacturerComponent;
  let fixture: ComponentFixture<EditAerospaceManufacturerComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        EditAerospaceManufacturerComponent
      ],
      providers: [
        AerospaceManufacturerService,
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

    fixture = TestBed.createComponent(EditAerospaceManufacturerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});