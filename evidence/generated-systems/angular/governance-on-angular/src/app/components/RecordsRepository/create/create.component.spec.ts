
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateRecordsRepositoryComponent } from './create.component';
import { RecordsRepositoryService } from '../../../services/RecordsRepository.service';
import { Router } from '@angular/router';

describe('CreateRecordsRepositoryComponent', () => {
  let component: CreateRecordsRepositoryComponent;
  let fixture: ComponentFixture<CreateRecordsRepositoryComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateRecordsRepositoryComponent
      ],
      providers: [
        RecordsRepositoryService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateRecordsRepositoryComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});