
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EditDataProcessingActivityComponent } from './edit.component';
import { DataProcessingActivityService } from '../../../services/DataProcessingActivity.service';

describe('EditDataProcessingActivityComponent', () => {
  let component: EditDataProcessingActivityComponent;
  let fixture: ComponentFixture<EditDataProcessingActivityComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        EditDataProcessingActivityComponent
      ],
      providers: [
        DataProcessingActivityService,
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

    fixture = TestBed.createComponent(EditDataProcessingActivityComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});