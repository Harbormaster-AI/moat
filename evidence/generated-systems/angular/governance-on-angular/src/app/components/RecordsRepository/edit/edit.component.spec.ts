
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EditRecordsRepositoryComponent } from './edit.component';
import { RecordsRepositoryService } from '../../../services/RecordsRepository.service';

describe('EditRecordsRepositoryComponent', () => {
  let component: EditRecordsRepositoryComponent;
  let fixture: ComponentFixture<EditRecordsRepositoryComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        EditRecordsRepositoryComponent
      ],
      providers: [
        RecordsRepositoryService,
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

    fixture = TestBed.createComponent(EditRecordsRepositoryComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});