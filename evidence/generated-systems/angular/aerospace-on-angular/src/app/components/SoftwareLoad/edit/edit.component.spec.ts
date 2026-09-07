
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EditSoftwareLoadComponent } from './edit.component';
import { SoftwareLoadService } from '../../../services/SoftwareLoad.service';

describe('EditSoftwareLoadComponent', () => {
  let component: EditSoftwareLoadComponent;
  let fixture: ComponentFixture<EditSoftwareLoadComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        EditSoftwareLoadComponent
      ],
      providers: [
        SoftwareLoadService,
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

    fixture = TestBed.createComponent(EditSoftwareLoadComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});