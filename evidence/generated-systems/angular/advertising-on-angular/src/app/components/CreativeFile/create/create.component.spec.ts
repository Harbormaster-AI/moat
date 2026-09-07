
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateCreativeFileComponent } from './create.component';
import { CreativeFileService } from '../../../services/CreativeFile.service';
import { Router } from '@angular/router';

describe('CreateCreativeFileComponent', () => {
  let component: CreateCreativeFileComponent;
  let fixture: ComponentFixture<CreateCreativeFileComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateCreativeFileComponent
      ],
      providers: [
        CreativeFileService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateCreativeFileComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});