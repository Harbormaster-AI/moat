
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EditSoftwareUpdateComponent } from './edit.component';
import { SoftwareUpdateService } from '../../../services/SoftwareUpdate.service';

describe('EditSoftwareUpdateComponent', () => {
  let component: EditSoftwareUpdateComponent;
  let fixture: ComponentFixture<EditSoftwareUpdateComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        EditSoftwareUpdateComponent
      ],
      providers: [
        SoftwareUpdateService,
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

    fixture = TestBed.createComponent(EditSoftwareUpdateComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});